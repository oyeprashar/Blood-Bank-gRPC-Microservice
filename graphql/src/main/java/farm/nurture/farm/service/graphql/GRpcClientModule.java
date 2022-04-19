package farm.nurture.farm.service.graphql;

import com.google.inject.AbstractModule;
import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import farm.nurture.farm.service.proto.BloodBankSystemServiceGrpc;

public class GRpcClientModule extends AbstractModule {

    private String grpcServiceAddress;

    public GRpcClientModule(String grpcServiceAddress) {
        this.grpcServiceAddress = grpcServiceAddress;
    }

    @Override
    protected void configure() {
        ManagedChannel channel = ManagedChannelBuilder.forTarget(grpcServiceAddress).usePlaintext().build();
        BloodBankSystemServiceGrpc.BloodBankSystemServiceFutureStub futureStub = BloodBankSystemServiceGrpc.newFutureStub(channel);
        bind(BloodBankSystemServiceGrpc.BloodBankSystemServiceFutureStub.class).toInstance(futureStub);
    }
}
