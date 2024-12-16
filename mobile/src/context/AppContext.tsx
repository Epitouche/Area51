// src/context/AppContext.tsx
import React, { createContext, useState, ReactNode, useEffect } from 'react';
import { AboutJson } from '../types';
import { checkToken } from '../service';

interface AppContextProps {
  serverIp: string;
  setServerIp: (ip: string) => void;
  aboutjson: AboutJson | undefined;
  setAboutJson: (aboutjson: AboutJson) => void;
  isConnected: boolean;
  setIsConnected: (isConnected: boolean) => void;
}

const AppContext = createContext<AppContextProps>({
  serverIp: '',
  setServerIp: () => { },
  aboutjson: undefined,
  setAboutJson: () => { },
  isConnected: false,
  setIsConnected: () => { },
});

interface AppProviderProps {
  children: ReactNode;
}

export default function AppProvider({ children }: AppProviderProps) {
  const [serverIp, setServerIp] = useState<string>('');
  const [aboutjson, setAboutJson] = useState<AboutJson>();
  const [isConnected, setIsConnected] = useState<boolean>(false);

  useEffect(() => {
    const checkConnection = async () => {
      if (await checkToken('token'))
        setIsConnected(true);
    };
    checkConnection();
  }, [serverIp]);

  return (
    <AppContext.Provider value={{ serverIp, setServerIp, aboutjson, setAboutJson, isConnected, setIsConnected }}>
      {children}
    </AppContext.Provider>
  );
}

export { AppContext };
