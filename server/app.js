new Vue({
    el: '#app',

    data: {
        ws: null, // Our websocket
        newMsg: '', // Holds new messages to be sent to the server
        chatContent: '', //updated list of chat messages displayed on screen
    },

    created: function() {
        var self = this;
        this.ws = new WebSocket('ws://' + window.location.host + '/ws');
        this.ws.addEventListener('message', function(e) {
            var msg = JSON.parse(e.data);
            self.chatContent += msg.message + '<br/>'; 
            var element = document.getElementById('chat-messages');
            element.scrollTop = element.scrollHeight; //Scroll down
        });
    },

    methods: {
        send: function () {
            if (this.newMsg != '') {
                this.ws.send(
                    JSON.stringify({
                        message: this.newMsg
                    }
                ));
                this.newMsg = ''; // Reset newMsg
            }
        },
    }
});