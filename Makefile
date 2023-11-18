RABBITMQ_USER=test
RABBITMQ_PASSWORD=secret
VHOST=events

setup-container:
	@echo Setup rabbitmq container
	@docker-compose up -d --quiet-pull

setup-rabbitmq:
	@echo -b setup user and password:
	@docker exec mq rabbitmqctl add_user ${RABBITMQ_USER} ${RABBITMQ_PASSWORD}
	# Give access
	@docker exec mq rabbitmqctl set_user_tags ${RABBITMQ_USER} administrator
	# Remove guest user
	@docker exec mq rabbitmqctl delete_user guest
	# Add vhost
	@docker exec mq rabbitmqctl add_vhost ${VHOST}
	# Give premission
	docker exec mq rabbitmqctl set_permissions -p ${VHOST} ${RABBITMQ_USER} ".*" ".*" ".*"

