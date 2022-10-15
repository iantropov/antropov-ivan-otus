import { Request } from 'express';
import { User } from '../user/user.enity';

interface RequestWithUser extends Request {
    user: User;
}

export default RequestWithUser;
