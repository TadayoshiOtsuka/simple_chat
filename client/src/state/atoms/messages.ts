import { Message } from "./../../models/message";
import { atom } from "recoil";

export const messageListAtom = atom<Message[]>({
  key: "messageList",
  default: [],
});
