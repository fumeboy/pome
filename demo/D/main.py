from socketserver import BaseRequestHandler, TCPServer

class EchoHandler(BaseRequestHandler):
    def handle(self):
        print('Got connection from', self.client_address)
        msg = self.request.recv(8192)
        if  msg:
            self.request.send(msg)
        

if __name__ == '__main__':
    print('start')
    serv = TCPServer(('', 8080), EchoHandler)
    serv.serve_forever()