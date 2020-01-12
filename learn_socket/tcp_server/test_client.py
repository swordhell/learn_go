#! /usr/bin/python3
# coding: utf-8
import asyncore, socket,struct
import time
import logging

logging.basicConfig(level=logging.DEBUG,
    format='%(asctime)s %(filename)s[line:%(lineno)d] %(levelname)s %(message)s',
    datefmt='%a, %d %b %Y %H:%M:%S')

msgstr = '''One night---it was on the twentieth of March,1888---I was returning
from a journey to a patient.
'''

class Client( asyncore.dispatcher ):

    def __init__( self, _host, _port ):
        asyncore.dispatcher.__init__( self )
        self.create_socket(socket.AF_INET, socket.SOCK_STREAM)
        self.connect( ( _host, _port ) )
        self.recv_buf_ = b''
        self.send_buf_ = b''
        self.post_data(msgstr)

    def post_data(self, _raw_buf):
        #
        # | short      | binary            |
        # | pack size  | encrypt binary    |
        # | binary data|                   |
        # small-endian unsigned int
        # 
        self.send_buf_ = self.send_buf_ + struct.pack("<h",len(_raw_buf))
        # can't concat str to bytes
        self.send_buf_ = self.send_buf_ + _raw_buf.encode()

    def handle_connect( self ):
        logging.debug("on handle_connect establishment")

    def handle_close( self ):
        self.close()
        logging.error("on handle_close socket cannt connect")

    def handle_read( self ):
        self.recv_buf_ += self.recv( 8192 )
        logging.debug('on handle_read read size: {0}'.format(len(self.recv_buf_)))
        #
        # | short      | binary            |
        # | pack size  | encrypt binary    |
        # | binary data|                   |
        # small-endian unsigned int
        # 
        while len(self.recv_buf_) > 2:
            if len(self.recv_buf_) < 2:
                return True
            stream_size = struct.unpack( "<h", self.recv_buf_[:2])[0]
            logging.debug("handle_read stream_size: {0}".format(stream_size))
            self.proc_data(self.recv_buf_[2:2+stream_size])
            self.recv_buf_ = self.recv_buf_[2+stream_size:]

    def writeable( self ):
        return ( len( self.send_buf_ ) > 0 )

    def handle_write( self ):
        if not self.writeable():
            time.sleep(1)
            return
        sent = self.send( self.send_buf_ )
        logging.info("handle_write, sent: %d"%sent)
        self.send_buf_ =  self.send_buf_[ sent: ]
        pass
    
    def proc_data(self,content):
        logging.debug("proc_data content: {}".format(content))

if __name__ == "__main__":
    logging.info("robot is launch")
    clients = []
    for i in range( 0, 100 ):
        clients.append( Client( '127.0.0.1', 8089) )
    asyncore.loop()
    pass
