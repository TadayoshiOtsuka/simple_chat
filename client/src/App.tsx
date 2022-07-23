import { useMessageList } from "./hooks/use_message_list";
import { useSubmitMessage } from "./hooks/use_submit_message";
import { useWebsocket } from "./hooks/use_websocket";

export const App = () => {
  useWebsocket();
  const messageList = useMessageList();
  const { input, setInput, submit } = useSubmitMessage();

  return (
    <div>
      <h1>Simple Chat</h1>
      <div>
        <input
          onChange={(e) => setInput(e.target.value)}
          value={input}
          placeholder="new message"
        />
        <button onSubmit={submit}>Submit</button>
      </div>
      <div>
        {messageList.map((m, i) => (
          <div key={i}>{m.content}</div>
        ))}
      </div>
    </div>
  );
};
