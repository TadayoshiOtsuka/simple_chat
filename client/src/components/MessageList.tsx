import { useMessageList } from "../hooks/use-message-list";

export const MessageList = () => {
  const messageList = useMessageList();

  return (
    <div>
      {messageList.map((m, i) => (
        <div key={i}>{m.content}</div>
      ))}
    </div>
  );
};
