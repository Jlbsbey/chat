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

export default function Register(props) {
    const [open, setOpen] = React.useState(false);
	const [registerDone, setRegisterDone] = React.useState(false);
    const [login, setLogin] = React.useState("");
	const [password, setPassword] = React.useState("");
    const [email, setEmail] = React.useState("");

    function loginChange(event) {
		setLogin(event.target.value);
	};
	function passwordChange(event) {
		setPassword(event.target.value);
	};
    function emailChange(event){
        setEmail(event.target.value);
    }

	function handleClickOpen() {
		setOpen(true);
	};
	function handleCancel() {
		setOpen(false);
	}

	function handleClose() {
		if (registerDone) {
			setOpen(false);
		}
	};


	function handleRegister() {
		let actn = {
			action: "register",
			object: "user",
			data: {
				username: login,
				password: password,
                email: email,
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
				alert("Error occured during registration");
			}

			return resp.json()
		}).then(data => {
			//console.log(data);
			if (data.success == true){
				alert("Registration successfull, now login");
			setRegisterDone(true);
			setOpen(false);
			} else if (data.success == false){
                alert(data.error_Message)
            }
		});
    }
    return (
		<>
			{/*<Button variant="standard" onClick={handleClickOpen}>
				Login
			</Button>*/}
			<Button variant="standard" onClick={handleClickOpen}>Register</Button>
			<Dialog open={open} onClose={handleClose}>
				<DialogTitle>Register</DialogTitle>
				<DialogContent>
					<DialogContentText>
						Enter your credentials
					</DialogContentText>
					<TextField
						autoFocus
						margin="dense"
						label="Login"
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
                    <TextField
						margin="dense"
						label="Email"
						type="email"
						fullWidth
						variant="standard"
						value={email}
						onChange={emailChange}
					/>
				</DialogContent>
				<DialogActions>
					<Button onClick={handleCancel}>Cancel</Button>
					<Button onClick={handleRegister}>Register</Button>
				</DialogActions>
			</Dialog>
		</>
	);
}

Register.propTypes = {
    backendIP: PropTypes.any.isRequired,
};
