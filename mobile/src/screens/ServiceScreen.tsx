import { Text, View } from 'react-native';
import { globalStyles } from '../styles/global_style';
import { AppContext } from '../context/AppContext';
import { useContext, useEffect } from 'react';
import { ServiceCard } from '../components';
import { getAboutJson } from '../service';

const ServiceButton = [
  {
    title: 'Github',
    image:
      'https://img.icons8.com/?size=100&id=3tC9EQumUAuq&format=png&color=000000',
    status: 'Connected',
  },
  {
    title: 'Google',
    image:
      'https://img.icons8.com/?size=100&id=3tC9EQumUAuq&format=png&color=000000',
    status: 'Disconnected',
  },
  {
    title: 'Youtube',
    image:
      'https://img.icons8.com/?size=100&id=3tC9EQumUAuq&format=png&color=000000',
    status: 'Connected',
  },
  {
    title: 'Ta darone',
    image:
      'https://img.icons8.com/?size=100&id=3tC9EQumUAuq&format=png&color=000000',
    status: 'Connected',
  },
  {
    title: 'Ton pere',
    image:
      'https://img.icons8.com/?size=100&id=3tC9EQumUAuq&format=png&color=000000',
    status: 'Disconnected',
  },
  {
    title: 'Github',
    image:
      'https://img.icons8.com/?size=100&id=3tC9EQumUAuq&format=png&color=000000',
    status: 'Connected',
  },
];

export default function ServiceScreen() {
  const { isBlackTheme, aboutjson, setAboutJson, serverIp } = useContext(AppContext);

  useEffect(() => {
    if (serverIp !== '') {
      getAboutJson(serverIp, setAboutJson);
    }
  }, [serverIp]);


  return (
    <View
      style={
        isBlackTheme ? globalStyles.wallpaperBlack : globalStyles.wallpaper
      }>
      <View style={globalStyles.container}>
        <Text
          style={isBlackTheme ? globalStyles.titleBlack : globalStyles.title}>
          Service Screen
        </Text>
        <View
          style={{
            width: '100%',
            flexWrap: 'wrap',
            flexDirection: 'row',
            justifyContent: 'center',
            gap: 10,
          }}>
          {aboutjson && aboutjson.server.services.map((service, index) => (
            <ServiceCard
              key={index}
              title={service.name}
              image={'https://img.icons8.com/?size=100&id=3tC9EQumUAuq&format=png&color=000000'}
              status={'Disconnected'}
            />
          ))}
        </View>
      </View>
    </View>
  );
}
