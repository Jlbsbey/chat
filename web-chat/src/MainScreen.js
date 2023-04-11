//React imports
import * as React from 'react';

//Material UI imports
import Box from '@mui/material/Box';
import Drawer from '@mui/material/Drawer';
import AppBar from '@mui/material/AppBar';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import Avatar from '@mui/material/Avatar';
import Menu from '@mui/material/Menu';
import MenuItem from '@mui/material/MenuItem';
import ListItemAvatar from '@mui/material/ListItemAvatar';

//Other imports

//Local imports
import RoomList from './Rooms/RoomList';
import ChatScreen from './Chat/ChatScreen';
import LoginDialog from './LoginDialog';
import AddRoom from './Rooms/AddRoom'

const drawerWidth = 0.2*window.innerWidth;
const backendIP = "http://localhost:8080"


let testRoom = {
    Name: "Test room 1",
    Messages: [
        
    ],
    ID: 1,
    //...
}

let testRoom2 = {
    Name: "Test room 2",
    Messages: [
        {Text: "Foo", Author: "TheUser1"},
        {Text: "Bar", Author: "TheUser2"},
        {Text: "FooBar", Author: "TheUser1"},
        {Text: "BarFoo", Author: "TheUser2"},
    ],
    ID: 2,
    //...
}

let testRoom3 = {
    Name: "Test room 3",
    Messages: [
        {Text: "Foo", Author: "TheUser1"},
        {Text: "Bar", Author: "TheUser2"},
        {Text: "FooBar", Author: "TheUser1"},
        {Text: "BarFoo", Author: "TheUser2"},
        {Text: "Foo", Author: "TheUser1"},
        {Text: "Bar", Author: "TheUser2"},
        {Text: "FooBar", Author: "TheUser1"},
        {Text: "BarFoo", Author: "TheUser2"},
        {Text: "Foo", Author: "TheUser1"},
        {Text: "Bar", Author: "TheUser2"},
        {Text: "FooBar", Author: "TheUser1"},
        {Text: "BarFoo", Author: "TheUser2"},
        {Text: "Foo", Author: "TheUser1"},
        {Text: "Bar", Author: "TheUser2"},
        {Text: "FooBar", Author: "TheUser1"},
        {Text: "BarFoo", Author: "TheUser2"},
        {Text: "Foo", Author: "TheUser1"},
        {Text: "Bar", Author: "TheUser2"},
        {Text: "FooBar", Author: "TheUser1"},
        {Text: "BarFoo", Author: "TheUser2"},
        {Text: "Foo", Author: "TheUser1"},
        {Text: "Bar", Author: "TheUser2"},
        {Text: "FooBar", Author: "TheUser1"},
        {Text: "BarFoo", Author: "TheUser2"},
        {Text: "Foo", Author: "TheUser1"},
        {Text: "Bar", Author: "TheUser2"},
        {Text: "FooBar", Author: "TheUser1"},
        {Text: "BarFoo", Author: "TheUser2"},
    ],
    ID: 3,
    //...
}

const emptyRoom = {
    Name: "",
    Messages: [],
    ID: 0,
    //...
}


export default function MainScreen() {
    //
    const [activeSession, setSession]=React.useState(0)
    const [userID, setUserID]=React.useState(0)
    const [userName, setUserName]=React.useState("")
    const [email, setEmail]=React.useState("")
    const [roomList, setRoomList] = React.useState([]);
    const [activeRoom, setActiveRoom] = React.useState(emptyRoom);


    function updateRoomList() {
        setRoomList([testRoom, testRoom2, testRoom3]);
        //place for fetch: action read room 
        //...
    }

    React.useEffect(() => {
        updateRoomList();
    });

    const [anchorEl, setAnchorEl] = React.useState(null);
    const open = Boolean(anchorEl);
    const handleClick = (event) => {
    event.preventDefault();
    setAnchorEl(event.currentTarget);
    };
    const handleClose = () => {
    setAnchorEl(null);
    };

    return (
        <Box sx={{ display: 'flex' }} height="100%">  {/*container for everything*/} 

            {/*AppBar is the blue bar with the title on top*/}
            <AppBar position="fixed" sx={{ zIndex: (theme) => theme.zIndex.drawer + 1 }}>
                <Toolbar>
                    
                    <Typography variant="h6" noWrap component="div" sx={{ flexGrow: 1 }}>
                        The Go Chat: {activeSession}, {userID}, {userName}, {email}
                    </Typography>
                        <ListItemAvatar onClick={handleClick} > 
                        <Avatar alt="User avatar" src="/folder/image.jpg" />
                    </ListItemAvatar>
                    
                    <Menu
                        id="demo-positioned-menu"
                        aria-labelledby="demo-positioned-button"
                        anchorEl={anchorEl}
                        open={open}
                        onClose={handleClose}
                        anchorOrigin={{
                        vertical: 'top',
                        horizontal: 'left',
                        }}
                        transformOrigin={{
                        vertical: 'top',
                        horizontal: 'left',
                        }}
                    >
                    <MenuItem onClick={handleClose}>Settings</MenuItem>
                    <LoginDialog backendIP={backendIP} session={activeSession} setSession={setSession} userID={userID} setUserID={setUserID} userName={userName} setUserName={setUserName} email={email} setEmail={setEmail}/>
                    <MenuItem onClick={handleClose}>Logout</MenuItem>
                    </Menu>
                </Toolbar>
            </AppBar>

            {/*Drawer is that thing on the left side*/}
            <Drawer
                variant="permanent"
                sx={{
                    width: drawerWidth,
                    flexShrink: 0,
                    [`& .MuiDrawer-paper`]: { width: drawerWidth, boxSizing: 'border-box' },
                }}
            >
                <Toolbar />
                <RoomList activeRoom={activeRoom} setActiveRoom={setActiveRoom} roomList={roomList}/>
                <AddRoom backendIP={backendIP} activeSession={activeSession} setSession={setSession} userID={userID} setUserID={setUserID}/>
            </Drawer>

            {/*This is the window with the chat*/}
            <ChatScreen activeRoom={activeRoom} setActiveRoom={setActiveRoom} activeSession={setSession} backendIP={backendIP}/>
        </Box>
    );
}
