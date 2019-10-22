import { Socket } from 'socket.io';
import { Events, IMessageData, IToken } from '../types';
import { verify } from '../promisified/jwt';
import { config, harmonyServer } from '..';
import { User } from '../schema/userSchema';
import { Message } from '../schema/messageSchema';

export default function onMessage(socket: Socket) {
  socket.on(Events.MESSAGE, (data: IMessageData) => {
    harmonyServer.Database.verifyToken(data.token)
      .then(userid => {
        User.findOne({ userid })
          .then(user => {
            if (user) {
              harmonyServer.Database.addMessage(
                userid,
                data.message,
                data.files
              );
              harmonyServer.getSocketServer().emit(Events.MESSAGE, {
                username: user.username,
                message: data.message,
                avatar: user.avatar,
                files: data.files
              });
            }
          })
          .catch(err => {
            console.log(err);
          });
      })
      .catch(() => {
        harmonyServer
          .getSocketServer()
          .emit(Events.INVALIDATE_SESSION, 'Invalid token');
      });
  });
}
