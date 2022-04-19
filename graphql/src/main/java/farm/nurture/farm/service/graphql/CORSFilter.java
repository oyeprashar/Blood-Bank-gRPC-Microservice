package farm.nurture.farm.service.graphql;

import io.netty.handler.codec.http.HttpHeaderNames;

import javax.ws.rs.HttpMethod;
import javax.ws.rs.container.ContainerRequestContext;
import javax.ws.rs.container.ContainerResponseContext;
import javax.ws.rs.container.ContainerResponseFilter;
import javax.ws.rs.ext.Provider;

@Provider
public class CORSFilter implements ContainerResponseFilter {
    @Override
    public void filter(ContainerRequestContext request, ContainerResponseContext response) {
        response.getHeaders().add(HttpHeaderNames.ACCESS_CONTROL_ALLOW_ORIGIN.toString(), "*");
        response.getHeaders().addAll(HttpHeaderNames.ACCESS_CONTROL_ALLOW_HEADERS.toString(),
                HttpHeaderNames.ORIGIN, HttpHeaderNames.CONTENT_TYPE, HttpHeaderNames.ACCEPT, HttpHeaderNames.AUTHORIZATION);
        response.getHeaders().add(HttpHeaderNames.ACCESS_CONTROL_ALLOW_CREDENTIALS.toString(), "true");
        response.getHeaders().addAll(HttpHeaderNames.ACCESS_CONTROL_ALLOW_METHODS.toString(),
                HttpMethod.GET, HttpMethod.POST, HttpMethod.PUT,
                HttpMethod.DELETE, HttpMethod.OPTIONS, HttpMethod.HEAD);
    }
}
