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
import Register from './Register';
import { updateRoomList } from './MainScreen';


export default function LoginDialog(props) {
	const [open, setOpen] = React.useState(true);
	const [loginDone, setLoginDone] = React.useState(false);
	const [login, setLogin] = React.useState("");
	const [password, setPassword] = React.useState("");

	function loginChange(event) {
		setLogin(event.target.value);
	};
	function passwordChange(event) {
		setPassword(event.target.value);
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
			object: "user",
			data: {
				username: login,
				password: password,
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
				props.setSession(data.session_id);
				props.setUserID(data.data.user_id);
				props.setUserName(data.data.username);
				if(props.email == ""){
					props.setEmail("-");
				} else if(props.email != ""){
					props.setEmail(data.data.email);
				}
			alert("Login successfull");
			setLoginDone(true);
			setOpen(false);
			}
		});
	}
	if(props.session == 0){
		return (
			<>
				<Button variant="standard" onClick={handleClickOpen}>
					Login
				</Button>
				<Dialog open={open} onClose={handleClose}>
					<DialogTitle>Login</DialogTitle>
					<DialogContent>
						<DialogContentText>
							Enter your credentials
						</DialogContentText>
						<TextField
							autoFocus
							margin="dense"
							label="Email address"
							type="email"
							fullWidth
							variant="standard"
							value={login}
							onChange={loginChange}
						/>
						<TextField
							margin="dense"
							label="Password"
							type="password"
							fullWidth
							variant="standard"
							value={password}
							onChange={passwordChange}
						/>
					</DialogContent>
					<DialogActions>
						<Register backendIP={props.backendIP} />
						<Button onClick={handleLogin}>Login</Button>
					</DialogActions>
				</Dialog>
			</>
		);
	}else{
	return (
		<>
			<Button variant="standard" onClick={handleClickOpen}>
				Login
			</Button>
			{/*<MenuItem onClick={handleClickOpen}>Login</MenuItem>*/}
			<Dialog open={open} onClose={handleClose}>
				<DialogTitle>Login</DialogTitle>
				<DialogContent>
					<DialogContentText>
						Enter your credentials
					</DialogContentText>
					<TextField
						autoFocus
						margin="dense"
						label="Email address"
						type="email"
						fullWidth
						variant="standard"
						value={login}
						onChange={loginChange}
					/>
					<TextField
						margin="dense"
						label="Password"
						type="password"
						fullWidth
						variant="standard"
						value={password}
						onChange={passwordChange}
					/>
				</DialogContent>
				<DialogActions>
					<Register backendIP={props.backendIP} />
					<Button onClick={handleCancel}>Cancel</Button>
					<Button onClick={handleLogin}>Login</Button>
				</DialogActions>
			</Dialog>
		</>
	);}
}

LoginDialog.propTypes = {
    backendIP: PropTypes.any.isRequired,
};