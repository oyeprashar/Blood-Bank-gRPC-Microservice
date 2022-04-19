package farm.nurture.farm.service.graphql;

import com.google.api.graphql.rejoiner.GrpcSchemaModule;
import com.google.api.graphql.rejoiner.Mutation;
import com.google.api.graphql.rejoiner.Query;
import com.google.api.graphql.rejoiner.SchemaModule;
import com.google.common.collect.ImmutableList;
import com.google.common.util.concurrent.SettableFuture;
import com.google.common.util.concurrent.ListenableFuture;
import com.google.inject.Inject;
import graphql.schema.DataFetchingEnvironment;
import graphql.schema.GraphQLFieldDefinition;

import farm.nurture.infra.util.StringUtils;
import java.util.List;


import farm.nurture.farm.service.proto.*;

/**
 * A Laminar GraphQL {@link SchemaModule} backed by a gRPC service.
 * https://code.nurture.farm/platform/laminargrpc/blob/master/src/main/proto/laminar.proto
 */
public class GraphqlSchemaModule extends GrpcSchemaModule {

    private static Status getErrorStatus ( String msg) {
        if ( null == msg) msg = "Security error.";
        return Status.newBuilder().setStatus(StatusCode.REQUEST_DENIED).
                addErrorMessages(msg).build();
    }

    @Inject
    BloodBankSystemServiceGrpc.BloodBankSystemServiceFutureStub futureClient;

    @Override
    protected void configureSchema() {

        /**
        //Select queries
        ImmutableList<GraphQLFieldDefinition> queryList =
                serviceToFields(BloodBankSystemServiceGrpc.BloodBankSystemServiceFutureStub.class,
                        ImmutableList.of("executeFindPassword", "executeFindBlood"));
        addQueryList(queryList);

        //IUD queries
        ImmutableList<GraphQLFieldDefinition> mutationList =
                serviceToFields(BloodBankSystemServiceGrpc.BloodBankSystemServiceFutureStub.class,
                        ImmutableList.of("executeAddUser", "executeAddUserBulk", "executeAddBlood", "executeAddBloodBulk", "execute"));
        addMutationList(mutationList);
        */
    }

    static String errMsg = "Security check failure";

    
    @Query("executeFindPassword")
    public ListenableFuture<FindPasswordResponse> FindPassword(
            DataFetchingEnvironment environment,
            FindPasswordRequest req) {

        boolean isVerified = FindPasswordOAuth(environment, req.getRequestHeaders());
        if  ( isVerified ) return futureClient.executeFindPassword(req);
        else {
            SettableFuture<FindPasswordResponse> errFuture = SettableFuture.create();
            errFuture.set(FindPasswordResponse.newBuilder().setStatus(getErrorStatus(errMsg)).build());
            return errFuture;
        }
    }

    protected boolean FindPasswordOAuth(DataFetchingEnvironment environment, RequestHeaders reqHeaders) {


        OAuthValidator.VerifyInput authAndAppToken = getVerifyInput(environment, reqHeaders);
        boolean checkPrivate = true;
        boolean isVerified = ! checkPrivate;
        if ( checkPrivate ) {
            OAuthValidator.Claim[] claims = new OAuthValidator.Claim[] {
                
            };
            try {
                String preferredUserName =  ( claims.length > 0  ) ?
                    OAuthValidator.getInstance().verify(authAndAppToken, claims) :
                    OAuthValidator.getInstance().verify(authAndAppToken);

                isVerified = ( preferredUserName.equals(reqHeaders.getPrefferedUserName()));
            } catch (Exception ex) { isVerified = false; errMsg = ex.getMessage();}
        }
        return isVerified;
    }

    @Mutation("executeAddUser")
    public ListenableFuture<AddUserResponse> AddUser(
            DataFetchingEnvironment environment,
            AddUserRequest req) {

        boolean isVerified = AddUserOAuth(environment, req.getRequestHeaders());
        if  ( isVerified ) return futureClient.executeAddUser(req);
        else {
            SettableFuture<AddUserResponse> errFuture = SettableFuture.create();
            errFuture.set(AddUserResponse.newBuilder().setStatus(getErrorStatus(errMsg)).build());
            return errFuture;
        }
    }

    protected boolean AddUserOAuth(DataFetchingEnvironment environment, RequestHeaders reqHeaders) {


        OAuthValidator.VerifyInput authAndAppToken = getVerifyInput(environment, reqHeaders);
        boolean checkPrivate = true;
        boolean isVerified = ! checkPrivate;
        if ( checkPrivate ) {
            OAuthValidator.Claim[] claims = new OAuthValidator.Claim[] {
                
            };
            try {
                String preferredUserName =  ( claims.length > 0  ) ?
                    OAuthValidator.getInstance().verify(authAndAppToken, claims) :
                    OAuthValidator.getInstance().verify(authAndAppToken);

                isVerified = ( preferredUserName.equals(reqHeaders.getPrefferedUserName()));
            } catch (Exception ex) { isVerified = false; errMsg = ex.getMessage();}
        }
        return isVerified;
    }

    @Mutation("executeAddUserBulk")
    public ListenableFuture<BulkAddUserResponse> AddUserBulk(
            DataFetchingEnvironment environment,
            BulkAddUserRequest req) {

        boolean isVerified = AddUserBulkOAuth(environment, req.getRequestHeaders());
        if  ( isVerified ) return futureClient.executeAddUserBulk(req);
        else {
            SettableFuture<BulkAddUserResponse> errFuture = SettableFuture.create();
            errFuture.set(BulkAddUserResponse.newBuilder().setStatus(getErrorStatus(errMsg)).build());
            return errFuture;
        }
    }

    protected boolean AddUserBulkOAuth(DataFetchingEnvironment environment, RequestHeaders reqHeaders) {


        OAuthValidator.VerifyInput authAndAppToken = getVerifyInput(environment, reqHeaders);
        boolean checkPrivate = true;
        boolean isVerified = ! checkPrivate;
        if ( checkPrivate ) {
            OAuthValidator.Claim[] claims = new OAuthValidator.Claim[] {
                
            };
            try {
                String preferredUserName =  ( claims.length > 0  ) ?
                    OAuthValidator.getInstance().verify(authAndAppToken, claims) :
                    OAuthValidator.getInstance().verify(authAndAppToken);

                isVerified = ( preferredUserName.equals(reqHeaders.getPrefferedUserName()));
            } catch (Exception ex) { isVerified = false; errMsg = ex.getMessage();}
        }
        return isVerified;
    }

    @Query("executeFindBlood")
    public ListenableFuture<FindBloodResponse> FindBlood(
            DataFetchingEnvironment environment,
            FindBloodRequest req) {

        boolean isVerified = FindBloodOAuth(environment, req.getRequestHeaders());
        if  ( isVerified ) return futureClient.executeFindBlood(req);
        else {
            SettableFuture<FindBloodResponse> errFuture = SettableFuture.create();
            errFuture.set(FindBloodResponse.newBuilder().setStatus(getErrorStatus(errMsg)).build());
            return errFuture;
        }
    }

    protected boolean FindBloodOAuth(DataFetchingEnvironment environment, RequestHeaders reqHeaders) {


        OAuthValidator.VerifyInput authAndAppToken = getVerifyInput(environment, reqHeaders);
        boolean checkPrivate = true;
        boolean isVerified = ! checkPrivate;
        if ( checkPrivate ) {
            OAuthValidator.Claim[] claims = new OAuthValidator.Claim[] {
                
            };
            try {
                String preferredUserName =  ( claims.length > 0  ) ?
                    OAuthValidator.getInstance().verify(authAndAppToken, claims) :
                    OAuthValidator.getInstance().verify(authAndAppToken);

                isVerified = ( preferredUserName.equals(reqHeaders.getPrefferedUserName()));
            } catch (Exception ex) { isVerified = false; errMsg = ex.getMessage();}
        }
        return isVerified;
    }

    @Mutation("executeAddBlood")
    public ListenableFuture<AddBloodResponse> AddBlood(
            DataFetchingEnvironment environment,
            AddBloodRequest req) {

        boolean isVerified = AddBloodOAuth(environment, req.getRequestHeaders());
        if  ( isVerified ) return futureClient.executeAddBlood(req);
        else {
            SettableFuture<AddBloodResponse> errFuture = SettableFuture.create();
            errFuture.set(AddBloodResponse.newBuilder().setStatus(getErrorStatus(errMsg)).build());
            return errFuture;
        }
    }

    protected boolean AddBloodOAuth(DataFetchingEnvironment environment, RequestHeaders reqHeaders) {


        OAuthValidator.VerifyInput authAndAppToken = getVerifyInput(environment, reqHeaders);
        boolean checkPrivate = true;
        boolean isVerified = ! checkPrivate;
        if ( checkPrivate ) {
            OAuthValidator.Claim[] claims = new OAuthValidator.Claim[] {
                
            };
            try {
                String preferredUserName =  ( claims.length > 0  ) ?
                    OAuthValidator.getInstance().verify(authAndAppToken, claims) :
                    OAuthValidator.getInstance().verify(authAndAppToken);

                isVerified = ( preferredUserName.equals(reqHeaders.getPrefferedUserName()));
            } catch (Exception ex) { isVerified = false; errMsg = ex.getMessage();}
        }
        return isVerified;
    }

    @Mutation("executeAddBloodBulk")
    public ListenableFuture<BulkAddBloodResponse> AddBloodBulk(
            DataFetchingEnvironment environment,
            BulkAddBloodRequest req) {

        boolean isVerified = AddBloodBulkOAuth(environment, req.getRequestHeaders());
        if  ( isVerified ) return futureClient.executeAddBloodBulk(req);
        else {
            SettableFuture<BulkAddBloodResponse> errFuture = SettableFuture.create();
            errFuture.set(BulkAddBloodResponse.newBuilder().setStatus(getErrorStatus(errMsg)).build());
            return errFuture;
        }
    }

    protected boolean AddBloodBulkOAuth(DataFetchingEnvironment environment, RequestHeaders reqHeaders) {


        OAuthValidator.VerifyInput authAndAppToken = getVerifyInput(environment, reqHeaders);
        boolean checkPrivate = true;
        boolean isVerified = ! checkPrivate;
        if ( checkPrivate ) {
            OAuthValidator.Claim[] claims = new OAuthValidator.Claim[] {
                
            };
            try {
                String preferredUserName =  ( claims.length > 0  ) ?
                    OAuthValidator.getInstance().verify(authAndAppToken, claims) :
                    OAuthValidator.getInstance().verify(authAndAppToken);

                isVerified = ( preferredUserName.equals(reqHeaders.getPrefferedUserName()));
            } catch (Exception ex) { isVerified = false; errMsg = ex.getMessage();}
        }
        return isVerified;
    }



    private OAuthValidator.VerifyInput getVerifyInput(DataFetchingEnvironment environment, RequestHeaders reqHeaders) {
        OAuthValidator.VerifyInput authAndAppToken = environment.getContext();
        if ( null == authAndAppToken.authToken) authAndAppToken.authToken = reqHeaders.getAuthToken();
        if ( null == authAndAppToken.appToken)  authAndAppToken.appToken = reqHeaders.getAppToken();
        return authAndAppToken;
    }
}
