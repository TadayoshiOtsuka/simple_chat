import { MessageInput } from "./components/message_input";
import { MessageList } from "./components/message_list";

export const App = () => {
  return (
    <div>
      <h1>Simple Chat</h1>
      <MessageInput />
      <MessageList />
    </div>
  );
};
