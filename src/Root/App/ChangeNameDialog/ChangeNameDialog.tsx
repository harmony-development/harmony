import React, { useState, useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { IAppState, IShowChangeNameDialog } from '../../../store/types';
import { toggleChangeNameDialog, updateUser } from '../../../store/actions/AppActions';
import { Dialog, DialogTitle, DialogContent, TextField, DialogActions, Button } from '@material-ui/core';
import UserIcon from './UserIcon/UserIcon';
import { socketServer } from '../../Root';
import { toast } from 'react-toastify';

const ChangeNameDialog: React.FC<{}> = () => {
  const dispatch = useDispatch();
  const [open, user] = useSelector((state: IAppState) => [state.nameDialog, state.user]);
  const [draftName, setDraftName] = useState(user.username);
  const [draftIcon, setDraftIcon] = useState(user.avatar);

  const setProfileAndClose = (): void => {
    if (localStorage.getItem('token') && typeof localStorage.getItem('token') === 'string') {
      socketServer.saveProfile({ username: draftName, avatar: draftIcon, token: localStorage.getItem('token') as string });
      dispatch(
        updateUser({
          username: draftName,
          avatar: draftIcon
        })
      );
    } else {
      toast.error('We were unable to save your profile.');
    }
    dispatch(toggleChangeNameDialog());
  };

  useEffect(() => {
    setDraftName(user.username);
    setDraftIcon(user.avatar);
  }, [open]);

  return (
    <Dialog open={open} onClose={(): IShowChangeNameDialog => dispatch(toggleChangeNameDialog())}>
      <DialogTitle>Change Username</DialogTitle>
      <DialogContent>
        <UserIcon setIcon={setDraftIcon} icon={draftIcon} />
      </DialogContent>
      <DialogContent>
        <TextField label='Username' value={draftName} onChange={(e: React.ChangeEvent<HTMLInputElement>): void => setDraftName(e.target.value)} />
      </DialogContent>
      <DialogActions>
        <Button onClick={setProfileAndClose} color='primary' autoFocus>
          Apply
        </Button>
      </DialogActions>
    </Dialog>
  );
};

export default ChangeNameDialog;
