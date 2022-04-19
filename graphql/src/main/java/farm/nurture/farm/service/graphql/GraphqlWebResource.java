package farm.nurture.farm.service.graphql;

import javax.ws.rs.GET;
import javax.ws.rs.Path;
import javax.ws.rs.Produces;
import javax.ws.rs.core.MediaType;
import javax.ws.rs.core.Response;

@Path("BloodBankSystemService/core/BloodBankSystemService/web")
@Produces(MediaType.TEXT_HTML)
public class GraphqlWebResource {

    @GET
    public Response getGraphiql() {
        ClassLoader classloader = Thread.currentThread().getContextClassLoader();
        return Response.ok(classloader.getResourceAsStream("index.html")).build();
    }
}
