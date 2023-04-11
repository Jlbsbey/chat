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

export default function NewRoom(props) {
	const [open, setOpen] = React.useState(false);
	const [loginDone, setLoginDone] = React.useState(false);
	const [name, setName] = React.useState("");

	function nameChange(event) {
		setName(event.target.value);
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
			action: "register",
			object: "room",
			data: {
				name: name,
				user_id: props.userID,
			},
		}

		//place for fetch: action login user
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
			//console.log(data);
			if (data.success == true){
				alert("Creating room was successfull");
			setLoginDone(true);
			setOpen(false);
			}
		});
	}

	return (
		<>
			{/*<Button variant="standard" onClick={handleClickOpen}>
				Login
			</Button>*/}
			<Button onClick={handleClickOpen}>New Room</Button>
			<Dialog open={open} onClose={handleClose}>
				<DialogTitle>New Room</DialogTitle>
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
				</DialogContent>
				<DialogActions>
                    <Button onClick={handleCancel}>Cancel</Button>
					<Button onClick={handleLogin}>Create</Button>
				</DialogActions>
			</Dialog>
		</>
	);
}

NewRoom.propTypes = {
    backendIP: PropTypes.any.isRequired,
};