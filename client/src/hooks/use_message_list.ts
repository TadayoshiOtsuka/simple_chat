import { websocketAtom } from "./../state/atoms/websocket";
import { useEffect } from "react";
import { messageListAtom } from "./../state/atoms/messages";
import { useRecoilState, useRecoilValue } from "recoil";
import { Message } from "../models/message";

export const useMessageList = (): Message[] => {
  const [messageList, setMessageList] = useRecoilState(messageListAtom);
  const socket = useRecoilValue(websocketAtom);

  useEffect(() => {
    socket.onmessage = (data) => {
      setMessageList((currVal) => [...currVal]);
    };
  }, [socket.onmessage]);

  return messageList;
};
