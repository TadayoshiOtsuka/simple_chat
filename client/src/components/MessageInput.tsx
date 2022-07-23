import { useSubmitMessage } from "../hooks/use-submit-message";

export const MessageInput = () => {
  const { input, setInput, submit } = useSubmitMessage();

  return (
    <div>
      <input
        onChange={(e) => setInput(e.target.value)}
        value={input}
        placeholder="new message"
      />
      <button onClick={submit}>Submit</button>
    </div>
  );
};
