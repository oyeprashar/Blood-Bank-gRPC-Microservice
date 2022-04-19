package farm.nurture.farm.service.graphql;

import com.google.gson.annotations.SerializedName;

import java.util.Collections;
import java.util.Map;

public class GraphqlRequest {

    @SerializedName("query")
    private String query;

    @SerializedName("operationName")
    private String operationName;

    @SerializedName("variables")
    private Map<String, Object> variables;

    public Map<String, Object> getVariables() {
        return null != variables ? variables : Collections.emptyMap();
    }

    public String getQuery() {
        return query;
    }

    public void setQuery(String query) {
        this.query = query;
    }

    public String getOperationName() {
        return operationName;
    }

    public void setOperationName(String operationName) {
        this.operationName = operationName;
    }

    public void setVariables(Map<String, Object> variables) {
        this.variables = variables;
    }
}
