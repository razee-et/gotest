FROM rmohr/activemq:5.15.9

ENV ACTIVEMQ_SSL_OPTS "-Djavax.net.ssl.keyStore=/opt/activemq/broker.ks -Djavax.net.ssl.keyStorePassword=password"

COPY activemq.xml /opt/activemq/conf/
RUN keytool -genkey -alias broker -dname "CN=localhost,OU=NGCSTATE,O=Acquia,L=Boston,ST=MA,C=US" -keyalg RSA \
  -ext "san=dns:localhost,dns:activemq,dns:activemq.activemq.svc.cluster.local" -keystore /opt/activemq/broker.ks -storepass password
