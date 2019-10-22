import { Socket } from 'socket.io';
import { Events, IGetUserData, IToken, IUser } from '../types';
import { harmonyServer } from '..';
import { User } from '../schema/userSchema';

export default function onGetUserData(socket: Socket) {
  socket.on(Events.GET_USER_DATA, (data: IGetUserData) => {
    harmonyServer.Database.verifyToken(data.token)
      .then(userid => {
        User.findOne({ userid })
          .then(user => {
            if (user) {
              harmonyServer.getSocketServer().emit(Events.GET_USER_DATA, {
                username: user.username,
                avatar: user.avatar,
                theme: user.theme
              });
            } else {
              socket.emit(Events.PROFILE_UPDATE_ERROR, 'You do not exist');
            }
          })
          .catch(err => {
            console.log(err);
          });
      })
      .catch(() => {
        socket.emit(Events.INVALIDATE_SESSION, 'Invalid token');
      });
  });
}
