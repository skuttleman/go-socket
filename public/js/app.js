console.log('here');

var socket = io();
socket.on('game move', function(response) {
  console.log(response);
});
