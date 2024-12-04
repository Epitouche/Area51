// src/context/AppContext.tsx
import React, { createContext, useState, ReactNode } from 'react';

interface AppContextProps {
  serverIp: string;
  setServerIp: (ip: string) => void;
}

const AppContext = createContext<AppContextProps>({
  serverIp: '',
  setServerIp: () => {},
});

interface AppProviderProps {
  children: ReactNode;
}

export function AppProvider({ children }: AppProviderProps) {
  const [serverIp, setServerIp] = useState<string>('');

  return (
    <AppContext.Provider value={{ serverIp, setServerIp }}>
      {children}
    </AppContext.Provider>
  );
}

export { AppContext };
export default AppProvider;
