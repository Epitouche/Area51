import { AboutJson } from './aboutJson';

export interface ServicesModalsProps {
  modalVisible: boolean;
  setModalVisible: (modalVisible: boolean) => void;
  services: AboutJson | undefined;
  setActionOrReaction: (actionOrReaction: ActionReaction) => void;
  isAction?: boolean;
}

export type ActionReaction = {
  id: number;
  name: string;
};


export type ConnectedService = {
  created_at: string;
  description: string;
  id: number;
  name: string;
  updated_at: string;
};