import { makeUniqueId } from '@apollo/client/utilities';
import { Message } from './types';

let messages = [];
let subscribers = [];

const DELETE_MESSAGE_TIMEOUT = 5000;

function notifySuscribers() {
    subscribers.forEach(callback => callback(messages));
}

function addMessage(text: string, type: Message['type']) {
    const newMessageId = makeUniqueId('message');
    messages = messages.concat({
        type,
        text,
        id: newMessageId
    });
    setTimeout(() => {
        removeMessage(newMessageId);
        notifySuscribers();
    }, DELETE_MESSAGE_TIMEOUT);
}

function removeMessage(messageIdToDelete: string) {
    messages = messages.filter(message => {
        return message.id !== messageIdToDelete;
    });
}

export const messageBroker = {
    subsribe: (callback: (messages: Message[]) => void) => {
        subscribers.push(callback);
    },
    addSuccessMessage(text: string) {
        addMessage(text, 'success');
        notifySuscribers();
    },
    addErrorMessage(text: string) {
        addMessage(text, 'error');
        notifySuscribers();
    }
};
