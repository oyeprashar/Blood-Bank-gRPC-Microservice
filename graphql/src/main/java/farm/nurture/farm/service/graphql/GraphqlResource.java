package farm.nurture.farm.service.graphql;

import com.google.inject.Inject;
import farm.nurture.infra.util.Logger;
import farm.nurture.infra.util.LoggerFactory;
import graphql.ExecutionInput;
import graphql.ExecutionResult;
import graphql.GraphQL;

import javax.ws.rs.Consumes;
import javax.ws.rs.POST;
import javax.ws.rs.Path;
import javax.ws.rs.Produces;
import javax.ws.rs.core.Context;
import javax.ws.rs.core.MediaType;
import java.util.List;
import java.util.Map;

@Path("BloodBankSystemService/core/BloodBankSystemService/graphql")
@Produces(MediaType.APPLICATION_JSON)
@Consumes(MediaType.APPLICATION_JSON)
public class GraphqlResource {

    static final Logger logger = LoggerFactory.getLogger(GraphqlResource.class);
    @Inject
    private GraphQL graphQL;

    @POST
    public Map<String, Object> graphqlHandler(@Context javax.ws.rs.core.HttpHeaders headers, GraphqlRequest request) {
        long startTime = System.currentTimeMillis();
        logger.debug("Request received start time : {}", startTime);

        OAuthValidator.VerifyInput verifyInput = new OAuthValidator.VerifyInput();
        List<String> authTokenL = headers.getRequestHeader("authtoken");
        verifyInput.authToken = ( null != authTokenL && authTokenL.size() == 1) ? authTokenL.get(0) : null;

        List<String> appTokenL = headers.getRequestHeader("apptoken");
        verifyInput.appToken = ( null != appTokenL && appTokenL.size() == 1) ? appTokenL.get(0) : null;

        List<String> ipL = headers.getRequestHeader("X-FORWARDED-FOR");
        verifyInput.ipAddress = ( null != ipL && ipL.size() == 1) ? ipL.get(0) : null;

        //create a graphql request execution object
        ExecutionInput executionInput =
                ExecutionInput.newExecutionInput()
                        .query(request.getQuery())
                        .context(verifyInput)
                        .operationName(request.getOperationName())
                        .variables(request.getVariables())
                        .build();

        //execute the request and send the response as a json
        ExecutionResult executionResult = graphQL.execute(executionInput);
        logger.debug("Request completed time taken {}", (System.currentTimeMillis() - startTime));
        return executionResult.toSpecification();
    }
}
