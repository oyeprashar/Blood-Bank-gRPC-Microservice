package farm.nurture.farm.service.graphql;

import com.auth0.jwt.JWT;
import com.auth0.jwt.JWTVerifier;
import com.auth0.jwt.algorithms.Algorithm;
import com.auth0.jwt.exceptions.TokenExpiredException;
import com.auth0.jwt.interfaces.DecodedJWT;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;
import farm.nurture.infra.util.Logger;
import farm.nurture.infra.util.LoggerFactory;
import farm.nurture.infra.util.StringUtils;
import org.bouncycastle.util.io.pem.PemObject;
import org.bouncycastle.util.io.pem.PemReader;

import java.io.File;
import java.io.FileNotFoundException;
import java.io.FileReader;
import java.io.IOException;
import java.security.KeyFactory;
import java.security.NoSuchAlgorithmException;
import java.security.PrivateKey;
import java.security.PublicKey;
import java.security.interfaces.RSAKey;
import java.security.spec.EncodedKeySpec;
import java.security.spec.InvalidKeySpecException;
import java.security.spec.PKCS8EncodedKeySpec;
import java.security.spec.X509EncodedKeySpec;
import java.util.Base64;
import java.util.Map;

public class OAuthValidator {

    public static class Claim {
        public String key;
        public String val;

        public Claim(String key, String val) {
            this.key = key;
            this.val = val;
        }
    }

    public static class VerifyInput {
        public String authToken = null;
        public String appToken = null;
        public String ipAddress = null;
    }

    private static final Logger logger = LoggerFactory.getLogger(OAuthValidator.class);
    private static OAuthValidator instance;

    public static OAuthValidator getInstance() {
        if ( null != instance) return instance;
        synchronized (OAuthValidator.class.getName()) {
            if ( null != instance) return instance;
            instance = new    OAuthValidator();
        }
        return instance;
    }

    JWTVerifier authVerifier = null;
    JWTVerifier appVerifier = null;

    /**
     * AUTH_KEY_PATH, APP_KEY_PATH is in k8 vault
     * The path will point to a S3 location
     * The file will be downloaded from this location at startup.
     * Once the file is downloaded, it can not be overwritten.
     */
    private OAuthValidator() {
        this(System.getenv("AUTH_KEY_PATH"),
                System.getenv("AUTH_KEY_ISSUER"),
                System.getenv("APP_KEY_PATH"),
                System.getenv("APP_KEY_ISSUER"));
    }

    private OAuthValidator(String authKeyPath, String authKeyIssuer,
                           String appKeyPath, String appKeyIssuer) {

        logger.info ( "Loading Auth verifier {} ", authKeyPath);
        this.authVerifier = loadKey (authKeyPath, authKeyIssuer);

        logger.info ( "Loading App verifier {} ", appKeyPath);
        this.appVerifier = loadKey (appKeyPath, appKeyIssuer);


    }

    private JWTVerifier loadKey (String aKeyPath, String aKeyIssuer) {

        if ( null == aKeyPath || null == aKeyIssuer) {
            logger.warn(
                "Key path or aKeyIssuer is missing. key = {}, issuer = {}", aKeyPath, aKeyIssuer);
            return null;
        }

        File publicKeyFile = new File(aKeyPath);
        if ( publicKeyFile.exists() ) {

            if ( publicKeyFile.canRead() ) {

                try {
                    Algorithm algorithm = Algorithm.RSA256((RSAKey)
                            readPublicKeyFromFile(aKeyPath, "RSA"));
                    logger.warn ("Checking the issuer {} for key {} ", aKeyIssuer, aKeyPath);
                    return JWT.require(algorithm).withIssuer(aKeyIssuer).build();

                } catch (IOException ex) {
                    logger.error( "Not able to load the public pem file from {}" , aKeyPath, ex);
                    System.exit(4);
                }

            } else {
                logger.error("Public key pem file is not readable from {}" , aKeyPath);
                System.exit(3);
            }
            return null;

        } else {
            logger.error("Public key pem file does not exist at {}" , aKeyPath);
            return null;
        }
    }

    /**
     * Signed by proper certificate
     * And the token is not expired.
     * @throws SecurityException
     */
    public String verify(VerifyInput input) throws SecurityException {
        if ( null == authVerifier) throw new SecurityException("System is not set properly. Auth Verifiers are missing.");
        if ( null == input.authToken) throw new SecurityException("Auth token is missing.");

        try {
            com.auth0.jwt.interfaces.DecodedJWT jwt = authVerifier.verify(input.authToken);

            String payloadStr = new String ( Base64.getDecoder().decode(jwt.getPayload()));
            JsonNode parent = new ObjectMapper().readTree(payloadStr);
            JsonNode preferredUserNameNode =  parent.path("preferred_username");


            String preferredUserName =  ( null == preferredUserNameNode) ? "" :  preferredUserNameNode.asText();

            if (StringUtils.isEmpty(preferredUserName)) {
                throw new SecurityException("Preferred user name is empty");
            }

            logger.debug("payload.preferredUserName: {}" , preferredUserName);
            return preferredUserName;

        } catch (TokenExpiredException ex) {
            logger.error("TokenExpiredException {}", ex.getMessage(), ex);
            throw new SecurityException("Security authorization failure.");

        } catch (Exception ex) {
            logger.error("Exception {}", ex.getMessage(), ex);
            throw new SecurityException("Security authorization failure.");

        }
    }

    public String verify(VerifyInput input, Claim... claims) throws SecurityException {

        String preferredUserName = verify(input);

        if ( null == appVerifier) throw new SecurityException("System is not set properly. App verifiers are missing.");
        if ( null == input.appToken) throw new SecurityException("App token is missing.");

        DecodedJWT appJwt =  appVerifier.verify(input.appToken);
        com.auth0.jwt.interfaces.Claim appPreferredClaim = appJwt.getClaim("preferred_username");
        String preferedUserNameApp = ( null == appPreferredClaim) ? "" : appPreferredClaim.asString();

        if ( preferredUserName.equals(preferedUserNameApp)) {
            if ( null != claims) {
                for ( Claim claim : claims) {
                    String value = appJwt.getClaim(claim.key ).asString();
                    if ( ! claim.val.equals(value)) {
                        throw new SecurityException("Claim " + claim.key + " is not accepted." );
                    }
                }
            }
        } else {
            throw new SecurityException("preferedUserNameApp did not match with app username token" + preferedUserNameApp );
        }
        return preferedUserNameApp;
    }

    private static byte[] parsePEMFile(File pemFile) throws IOException {
        if (!pemFile.isFile() || !pemFile.exists()) {
            throw new FileNotFoundException(String.format("The file '%s' doesn't exist.", pemFile.getAbsolutePath()));
        }
        PemReader reader = new PemReader(new FileReader(pemFile));
        PemObject pemObject = reader.readPemObject();
        byte[] content = pemObject.getContent();
        reader.close();
        return content;
    }

    private static PublicKey getPublicKey(byte[] keyBytes, String algorithm) {
        PublicKey publicKey = null;
        try {
            KeyFactory kf = KeyFactory.getInstance(algorithm);
            EncodedKeySpec keySpec = new X509EncodedKeySpec(keyBytes);
            publicKey = kf.generatePublic(keySpec);
        } catch (NoSuchAlgorithmException e) {
            System.out.println("Could not reconstruct the public key, the given algorithm could not be found.");
        } catch (InvalidKeySpecException e) {
            System.out.println("Could not reconstruct the public key");
        }

        return publicKey;
    }

    private static PrivateKey getPrivateKey(byte[] keyBytes, String algorithm) {
        PrivateKey privateKey = null;
        try {
            KeyFactory kf = KeyFactory.getInstance(algorithm);
            EncodedKeySpec keySpec = new PKCS8EncodedKeySpec(keyBytes);
            privateKey = kf.generatePrivate(keySpec);
        } catch (NoSuchAlgorithmException e) {
            System.out.println("Could not reconstruct the private key, the given algorithm could not be found.");
        } catch (InvalidKeySpecException e) {
            System.out.println("Could not reconstruct the private key");
        }

        return privateKey;
    }

    private static PublicKey readPublicKeyFromFile(String filepath, String algorithm) throws IOException {
        logger.info ( "Reading readPublicKeyFromFile  {}", filepath);
        byte[] bytes = parsePEMFile(new File(filepath));
        return getPublicKey(bytes, algorithm);
    }

    public static PrivateKey readPrivateKeyFromFile(String filepath, String algorithm) throws IOException {
        byte[] bytes = parsePEMFile(new File(filepath));
        return getPrivateKey(bytes, algorithm);
    }

    public static void main(String[] args) throws Exception {
        try {
            OAuthValidator oauth = OAuthValidator.getInstance();
        } catch (SecurityException ex) {
            ex.printStackTrace();
        }
    }

}
