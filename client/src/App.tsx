import { Button, CssBaseline, Input } from "@mui/material";
import { useMessageList } from "./hooks/use_message_list";
import { useSubmitMessage } from "./hooks/use_submit_message";

export const App = () => {
  const messageList = useMessageList();
  const { input, setInput, submit } = useSubmitMessage();

  return (
    <div>
      <h1>Simple Chat</h1>
      <div>
        <Input
          onChange={(e) => setInput(e.target.value)}
          value={input}
          placeholder="new message"
        />
        <Button variant="contained" onClick={submit}>
          Submit
        </Button>
      </div>
      <div>
        {messageList.map((m, i) => (
          <div key={i}>{m.content}</div>
        ))}
      </div>
    </div>
  );
};
