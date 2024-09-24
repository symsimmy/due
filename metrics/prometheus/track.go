package prometheus

const (
	GateServerReceiveClientMessageCount     = "gate_server_receive_client_message_count"
	GateServerReceiveClientMessageBytes     = "gate_server_receive_client_message_bytes"
	GateServerSendToClientMessageCount      = "gate_server_send_to_client_message_count"
	GateServerSendToClientMessageBytes      = "gate_server_send_to_client_message_bytes"
	GateServerSendToServerMessageCount      = "gate_server_send_to_server_message_count"
	GateServerReceiveFromServerMessageCount = "gate_server_receive_from_server_message_count"
	GateServerClientReconnectCount          = "gate_server_client_reconnect_count"
	GateServerTotalOnlinePlayer             = "gate_server_total_online_player"
	GateServerWriteError                    = "gate_server_write_error"
	ServerInternalWriteError                = "server_internal_write_error"
	ServerReceiveHandleError                = "server_receive_handle_error"
	ServerRpcWriteDuration                  = "server_rpc_write_duration"
	ServerRpcHandleDuration                 = "server_rpc_handle_duration"
	GateServerWriteDuration                 = "gate_server_write_duration"
)
