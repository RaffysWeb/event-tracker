FROM confluentinc/cp-server-connect-base:7.4.1
ENV CONNECT_PLUGIN_PATH="/opt/kafka/plugins"
RUN confluent-hub install --no-prompt confluentinc/kafka-connect-cassandra:2.0.5