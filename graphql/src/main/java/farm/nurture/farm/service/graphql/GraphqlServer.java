package farm.nurture.farm.service.graphql;

import com.fasterxml.jackson.annotation.JsonAutoDetect;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.PropertyAccessor;
import com.fasterxml.jackson.databind.DeserializationFeature;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.SerializationFeature;
import com.fasterxml.jackson.jaxrs.json.JacksonJaxbJsonProvider;
import com.google.api.graphql.rejoiner.SchemaProviderModule;
import com.google.inject.Guice;
import com.google.inject.Injector;
import farm.nurture.infra.util.Logger;
import farm.nurture.infra.util.LoggerFactory;
import farm.nurture.infra.util.StringUtils;
import io.netty.channel.Channel;
import org.glassfish.jersey.netty.httpserver.NettyHttpContainerProvider;
import org.glassfish.jersey.server.ResourceConfig;

import java.net.URI;

/**
 * This is graphql server implementation for laminar grpc integration.
 *
 * <p>Client can interact with this server using below endpoint `/platform/laminar/graphql`
 *
 * <p>A web frontend for viewing complete schema and running queries can use below endpoint
 * `/platform/laminar/web`
 */
public class GraphqlServer {

    private static final Logger logger = LoggerFactory.getLogger(GraphqlServer.class);

    private static void startLaminarGraphqlServer(int graphqlServerPort, String grpcServiceAddress) throws InterruptedException {
        logger.info("Starting graphql laminar server");
        long start = System.currentTimeMillis();

        // Create injector modules that is used across captain server
        Injector parentInjector =
                Guice.createInjector(
                        new SchemaProviderModule(), // Guice module that provides the generated GraphQLSchem
                        new GRpcClientModule(grpcServiceAddress), // grpc service stub provider
                        new GraphqlSchemaModule() //Schema for graphql module for laminar
                );

        Injector childInjector = Guice.createInjector(new GraphqlResourceModule(parentInjector));

        // start captain server
        startNettyServer(childInjector, graphqlServerPort, start);
    }

    private static void startNettyServer(Injector injector, int graphqlServerPort, long start) {
        URI uri = URI.create("http://0.0.0.0:" + graphqlServerPort + "/");
        GraphqlWebResource webResource = injector.getInstance(GraphqlWebResource.class);
        GraphqlResource graphqlResource = injector.getInstance(GraphqlResource.class);
        ResourceConfig resourceConfig = new ResourceConfig();
        resourceConfig.register(new CORSFilter());
        resourceConfig.register(getJacksonProvider()).register(webResource);
        resourceConfig.register(getJacksonProvider()).register(graphqlResource);

        final Channel server = NettyHttpContainerProvider.createHttp2Server(uri, resourceConfig, null);
        server.flush();
        Runtime.getRuntime().addShutdownHook(new Thread(server::close));
        logger.info("Laminar graphql server Started in {} ms @ {}", (System.currentTimeMillis() - start), uri);
    }

    private static JacksonJaxbJsonProvider getJacksonProvider() {
        ObjectMapper mapper = new ObjectMapper();
        mapper.configure(DeserializationFeature.FAIL_ON_UNKNOWN_PROPERTIES, false);
        mapper.setSerializationInclusion(JsonInclude.Include.ALWAYS);
        mapper.setVisibility(PropertyAccessor.ALL, JsonAutoDetect.Visibility.ANY);
        mapper.enable(SerializationFeature.INDENT_OUTPUT);

        JacksonJaxbJsonProvider provider = new JacksonJaxbJsonProvider();
        provider.setMapper(mapper);
        return provider;
    }

    /** Main entry point for this server. */
    public static void main(String[] args) throws InterruptedException {
        String grpcServiceAddress = args[0];
        if(StringUtils.isEmpty(grpcServiceAddress)) {
            logger.error("grpcServiceAddress is not given");
            System.exit(1);
        }

        int graphqlServerPort = 9090;
        if(StringUtils.isNonEmpty(args[1])) {
            graphqlServerPort = Integer.parseInt(args[1]);
        }

        startLaminarGraphqlServer(graphqlServerPort, grpcServiceAddress);
    }
}
