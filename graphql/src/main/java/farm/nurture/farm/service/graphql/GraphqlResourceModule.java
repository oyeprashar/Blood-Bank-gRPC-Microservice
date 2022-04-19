package farm.nurture.farm.service.graphql;

import com.google.api.graphql.execution.GuavaListenableFutureSupport;
import com.google.api.graphql.rejoiner.Schema;
import com.google.inject.AbstractModule;
import com.google.inject.Injector;
import com.google.inject.Key;
import farm.nurture.infra.util.Logger;
import farm.nurture.infra.util.LoggerFactory;
import graphql.GraphQL;
import graphql.execution.instrumentation.ChainedInstrumentation;
import graphql.execution.instrumentation.Instrumentation;
import graphql.schema.GraphQLSchema;

import java.util.ArrayList;
import java.util.List;

public final class GraphqlResourceModule extends AbstractModule {

    private static final Logger logger = LoggerFactory.getLogger(GraphqlResourceModule.class);

    private final Injector parentInjector;

    public GraphqlResourceModule(Injector parentInjector) {
        this.parentInjector = parentInjector;
    }

    @Override
    protected void configure() {
        // bind graphql schema to graphql execution unit
        bindGraphql();

        // bind resources
        bind(GraphqlResource.class).toInstance(new GraphqlResource());
        bind(GraphqlWebResource.class).toInstance(new GraphqlWebResource());
    }

    void bindGraphql() {
        logger.info("Initializing graphql schema....");
        long start = System.currentTimeMillis();
        // Get a graphql schema
        GraphQLSchema schema =
                parentInjector.getProvider(Key.get(GraphQLSchema.class, Schema.class)).get();

        //setup graphql
        GraphQL graphQL = setUpGraphQl(schema);
        bind(GraphQL.class).toInstance(graphQL);
        logger.info(
                "Graphql schema initialization successful in {} ms", (System.currentTimeMillis() - start));
    }

    private GraphQL setUpGraphQl(GraphQLSchema schema) {
        /*
         * Add instrumentation for listenable future which is used for calling grpc services.
         * Also add prometheus instrumentation for field fetch in graphql query
         */
        List<Instrumentation> instrumentations = new ArrayList<>();
        instrumentations.add(GuavaListenableFutureSupport.listenableFutureInstrumentation());
        //instrumentations.add(new TracingInstrumentation());
        //add instrumentations
        Instrumentation instrumentation = new ChainedInstrumentation(instrumentations);
        // Create a new graphql schema.
        return GraphQL.newGraphQL(schema).instrumentation(instrumentation).build();
    }
}
