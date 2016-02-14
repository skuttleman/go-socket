console.log('here');

var socket = io();
socket.on('chat message', function(response) {
  console.log(response);
});
