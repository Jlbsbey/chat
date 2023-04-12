//React imports
import * as React from 'react';

//Material UI imports
import ListItem from '@mui/material/ListItem';
import ListItemText from '@mui/material/ListItemText';
import Typography from '@mui/material/Typography';
import ListItemAvatar from '@mui/material/ListItemAvatar';
import Avatar from '@mui/material/Avatar';
import Menu from '@mui/material/Menu';
import MenuItem from '@mui/material/MenuItem';

//Other imports
import PropTypes from 'prop-types';
import { ListItemButton } from '@mui/material';

//Local imports
export default function MessageListItem(props) {
    const [anchorEl, setAnchorEl] = React.useState(null);
    const open = Boolean(anchorEl);
    const handleClick = (event) => {
    event.preventDefault();
    setAnchorEl(event.currentTarget);
    };
    const handleClose = () => {
    setAnchorEl(null);
    };

    MessageListItem.propTypes = {
        Message: PropTypes.any.isRequired,
    };
    return (
        <ListItem key={props.Message.message_id}>
            <ListItemButton >
                <ListItemAvatar onContextMenu={handleClick} > 
                    <Avatar alt="User avatar" src="/folder/image.jpg" />
                </ListItemAvatar>
                <ListItemText 
                disableTypography
                primary={<Typography sx={{color: '#8888FF'}}> {props.Message.author} </Typography>} 
                secondary={<Typography> {props.Message.Content.text} </Typography>} 
                />
                </ListItemButton>
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
                <MenuItem onClick={handleClose}>User's account</MenuItem>
                <MenuItem onClick={handleClose}>Delete message</MenuItem>
                <MenuItem onClick={handleClose}>Mark message</MenuItem>
                </Menu>  
        </ListItem>
    )
}
