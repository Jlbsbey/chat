//React imports
import * as React from 'react';

//Material UI imports
import Button from '@mui/material/Button';
import TextField from '@mui/material/TextField';
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import DialogContentText from '@mui/material/DialogContentText';
import DialogTitle from '@mui/material/DialogTitle';

//Other imports
import PropTypes from 'prop-types';
import MenuItem from '@mui/material/MenuItem';

//Local imports
import NewRoom from './NewRoom'


export default function AddRoom(props) {
	const [open, setOpen] = React.useState(false);
	const [loginDone, setLoginDone] = React.useState(false);
	const [name, setName] = React.useState("");
	const [inv_code, setCode] = React.useState(0);

	function nameChange(event) {
		setName(event.target.value);
	};
	function codeChange(event) {
		setCode(event.target.value);
	};

	function handleClickOpen() {
		setOpen(true);
	};
	function handleCancel() {
		setOpen(false);
	}

	function handleClose() {
		if (loginDone) {
			setOpen(false);
		}
	};

	function handleLogin() {
		let actn = {
			action: "login",
			object: "room",
			data: {
				name: name,
				invite_code: inv_code,
				user_id: props.userID,
			},
		}
		//add to room
		fetch(props.backendIP.concat("/"), {
			method: 'POST', 
			mode: 'cors', 
			cache: 'no-cache', 
			credentials: 'same-origin', 
			headers: {
			  	'Content-Type': 'application/json'
			},
			redirect: 'follow', 
			referrerPolicy: 'no-referrer', 
			body: JSON.stringify(actn),
		}).then(resp => {
			//The place where you should check if request was successfull and read info about response like headers
			if (!resp.ok) {
				alert("Error occured during login");
			}

			return resp.json()
		}).then(data => {
			if (data.success == true){
				alert("Adding room was successfull");
			//setLoginDone(true);
			setOpen(true);
			}
		});
	}

	return (
		<>
			{/*<Button variant="standard" onClick={handleClickOpen}>
				Login
			</Button>*/}
			<Button onClick={handleClickOpen}>+ New Room</Button>
			<Dialog open={open} onClose={handleClose}>
				<DialogTitle>Add Room</DialogTitle>
				<DialogContent>
					<DialogContentText>
						Enter your credentials
					</DialogContentText>
					<TextField
						autoFocus
						margin="dense"
						label="Room name"
						type="email"
						fullWidth
						variant="standard"
						value={name}
						onChange={nameChange}
					/>
					<TextField
						margin="dense"
						label="Invite code"
						type="number"
						fullWidth
						variant="standard"
						value={inv_code}
						onChange={codeChange}
					/>
				</DialogContent>
				<DialogActions>
					<NewRoom backendIP={props.backendIP}/>
					<Button onClick={handleCancel}>Cancel</Button>
					<Button onClick={handleLogin}>Add</Button>
				</DialogActions>
			</Dialog>
		</>
	);
}

AddRoom.propTypes = {
    backendIP: PropTypes.any.isRequired,
};