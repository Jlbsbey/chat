//React imports
import * as React from 'react';

//Material UI imports
import ListItem from '@mui/material/ListItem';
import ListItemButton from '@mui/material/ListItemButton';
import ListItemText from '@mui/material/ListItemText';

//Other imports
import PropTypes from 'prop-types';

//Local imports

export default function RoomListItem(props) {
    function changeRoom() {
        props.setActiveRoom(props.Room);
        let readmsg={
            action: "read",
            object: "message",
            data: {
                room_id: props.activeRoom.ID,
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
            }
            //setUserText("a")
            //setUserText("");
            
        });
    }

    return (
        <ListItem key={props.Room.ID} disablePadding>
            <ListItemButton onClick={changeRoom} selected={props.Room.ID===props.activeRoom.ID}>
                <ListItemText primary={props.Room.Name} />
            </ListItemButton>
        </ListItem>
    )
};

RoomListItem.propTypes = {
    Room: PropTypes.any.isRequired,
    activeRoom: PropTypes.any.isRequired,
    setActiveRoom: PropTypes.any.isRequired,
};