//React imports
import * as React from 'react';

//Material UI imports
import List from '@mui/material/List';
import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import TextField from '@mui/material/TextField';
import Button from '@mui/material/Button';
import SendIcon from '@mui/icons-material/Send';
import Paper from '@mui/material/Paper';

//Other imports
import PropTypes from 'prop-types';

//Local imports
import MessageListItem from './MessageListItem';

export default function ChatScreen(props) {
    const [messageAuthor, setAuthor] = React.useState("")
    const [userText, setUserText] = React.useState("");

    const [currentTime, setCurrentTime] = React.useState(new Date());
    const ws = React.useRef(null)
    function userTextChange(event) {
        setUserText(event.target.value);
    }

    React.useEffect(() => {
        ws.current = new WebSocket("ws://localhost:8080/ws");
        ws.current.onopen = () => console.log("");
        ws.current.onclose = () => console.log("");
        const wsCurrent = ws.current;

        return () => {
            //wsCurrent.close();
        };
    }, []);

    React.useEffect(()=>{
        if(!ws.current) return;

        ws.current.onmessage = e => {
            const message = JSON.parse(e.data)
            receiveMessage(message)
        };
    }, []);

    function receiveMessage(message){
        let msg={
            author_id: "",
            content:{
                text: "",
            },
            date: Date,
            is_forwarded: false, //placeholder
            reply_message_id: 0, //placeholder
            author: "",
            }
        if(message.action == "sync" && message.object == "message"){    
            msg.author_id = message.author_id;
            msg.content.text = message.content.text;
            msg.date = message.date;
            msg.author = message.author;
            msg.is_forwarded=false;
            msg.reply_message_id=0;
            props.activeRoom.Messages.push(msg);
        }
    props.activeRoom.Messages.push(msg);
    }
    
    function sendMessage(event) {
        setCurrentTime(new Date());
        let actn = {
            action: "register",
            object: "message",
            data: {
                room_id: props.activeRoom.ID,
                author_id: props.userID,
                content:{
                    text: userText,
                },
                date: currentTime,
                is_forwarded: false, //placeholder
                reply_message_id: 0, //placeholder


            }        
        }
        let msg={
            author_id: "",
            content:{
                text: "",
            },
            date: Date,
            is_forwarded: false, //placeholder
            reply_message_id: 0, //placeholder
            author: "",
        }
        msg.author_id = props.userID;
        msg.content.text = userText;
        msg.date = currentTime;
        msg.author = props.username;
        msg.is_forwarded=false;
        msg.reply_message_id=0;
        props.activeRoom.Messages.push(msg);
        
        props.setActiveRoom(props.activeRoom)
        ws.current.send(JSON.stringify(actn))
        setUserText("");     
        setAuthor(props.username);   
        //отправить сообщение на сервер и загрузить сообщение обратно С АВТОРОМ
        /*fetch(props.backendIP.concat("/"), {
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
			let readmsg={
                action: "read",
                object: "message",
                data: {
                    message_id: data.data.message_id,
                }
            }

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
                body: JSON.stringify(readmsg),
            }).then(resp => {
                //The place where you should check if request was successfull and read info about response like headers
                if (!resp.ok) {
                    alert("Error occured during login");
                }
                return resp.json()
            }).then(data => {
                let msg={
                    author_id: "",
                    content:{
                        text: "",
                    },
                    date: Date,
                    is_forwarded: false, //placeholder
                    reply_message_id: 0, //placeholder
                    author: "",
                    }
                for(let i=0; i< data.data.length; i++){
                    msg.author_id = data.data[i].author_id;
                    msg.content.text = data.data[i].content.text;
                    msg.date = data.data[i].date;
                    msg.author = data.data[i].author;
                    msg.is_forwarded=false;
                    msg.reply_message_id=0;
                    props.activeRoom.Messages.push(msg);
                    props.setActiveRoom(props.activeRoom)
                    //setRoomList(roomList => [...roomList, room])
                
                }
                setUserText("");
                
            });
            
		});*/

        
    }
    return (
        <Box m="10" sx={{ flexGrow: 1, pl: "5%", pr: "5%"}} key={props.activeRoom}>
            <Toolbar />
            <Paper elevation={3} sx={{mt:"1%", mb:"1%"}}> {/*List with messages*/}
                <List> 
                    {props.activeRoom.Messages.map((message, index) => (
                        <MessageListItem Message={message}/>
                    ))}
                </List>
            </Paper>

            <Paper elevation={3} sx={{ top: 'auto', bottom: 0, mb:"1%", position:"sticky"}}> {/*Toolbar with text field and send button*/}
                <Toolbar>
                    <TextField label="Type message: " variant="standard" fullWidth autoFocus sx={{mr:"2%"}} value={userText} onChange={userTextChange}/>
                    <Button variant="contained" endIcon={<SendIcon />} onClick={sendMessage}>
                        Send
                    </Button>
                </Toolbar>
            </Paper>
        </Box>
    )
};

ChatScreen.propTypes = {
    activeRoom: PropTypes.any.isRequired,
    setActiveRoom: PropTypes.any.isRequired,
};