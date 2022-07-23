import { useState } from "react";
import { websocketAtom } from "../state/atoms/websocket";
import { useRecoilValue } from "recoil";

export const useSubmitMessage = () => {
  const socket = useRecoilValue(websocketAtom);
  const [input, setInput] = useState<string>("");

  const submit = () => {
    if (input.length === 0) return;
    socket.send(JSON.stringify(input));
    setInput("");
  };

  return { input, setInput, submit };
};
