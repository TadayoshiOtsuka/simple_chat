import { websocketAtom } from "./../state/atoms/websocket";
import { useEffect } from "react";
import { messageListAtom } from "./../state/atoms/messages";
import { useRecoilState, useRecoilValue } from "recoil";
import { Message } from "../models/message";

export const useMessageList = (): Message[] => {
  const socket = useRecoilValue(websocketAtom);
  const [messageList, setMessageList] = useRecoilState(messageListAtom);

  useEffect(() => {
    socket.onmessage = (msg) => {
      const content = JSON.parse(msg.data as string);
      setMessageList((currVal) => [...currVal, { content: content }]);
    };
  }, [socket.onmessage]);

  return messageList;
};
