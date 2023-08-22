const connectLiveEventsWebSocket = (params) => {
  // const socket = new WebSocket("ws://localhost:4001/live-events");
  socket.onopen = () => {
    console.log('socket open')
  };
  return socket;
};

const subscribeToData = (socket, onDataReceived) => {
  socket.onmessage(newData => {
    console.log('newData')
    console.log(newData)
    onDataReceived(newData);
  });
};

const disconnectLiveEventsWebSocket = socket => {
  socket.disconnect();
};

export { connectLiveEventsWebSocket, subscribeToData, disconnectLiveEventsWebSocket };
