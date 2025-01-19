import { Text, View } from 'react-native';
import { globalStyles } from '../styles/global_style';
import { AppContext } from '../context/AppContext';
import { useContext, useEffect, useState } from 'react';
import { DeconnectionPopUp, ServiceCard } from '../components';
import { getToken, refreshServices } from '../service';

export default function ServiceScreen() {
  const {
    isBlackTheme,
    servicesConnected,
    aboutJson,
    serverIp,
    setServicesConnected,
    setAboutJson,
  } = useContext(AppContext);
  const [token, setToken] = useState('');

  const [needRefresh, setNeedRefresh] = useState(false);
  const [modalVisible, setModalVisible] = useState(false);
  const [selectedServices, setSelectedService] = useState('');

  useEffect(() => {
    const getMyToken = async () => {
      await getToken('token', setToken);
    };
    getMyToken();
    refreshServices({
      serverIp,
      setAboutJson,
      setServicesConnected,
      aboutJson,
    });
  }, [serverIp]);

  useEffect(() => {
    if (needRefresh) {
      setTimeout(() => {
        refreshServices({
          serverIp,
          setAboutJson,
          setServicesConnected,
          aboutJson,
        });
        setNeedRefresh(false);
      }, 300);
    }
  }, [needRefresh]);

  return (
    <View
      style={
        isBlackTheme ? globalStyles.wallpaperBlack : globalStyles.wallpaper
      }>
      <View style={globalStyles.container}>
        <Text
          style={isBlackTheme ? globalStyles.titleBlack : globalStyles.title}
          accessibilityLabel="Services">
          Services
        </Text>
        <View
          style={{
            width: '100%',
            flexWrap: 'wrap',
            flexDirection: 'row',
            justifyContent: 'center',
            gap: 10,
          }}>
          {servicesConnected &&
            aboutJson &&
            servicesConnected.services &&
            servicesConnected.services.map((service, index) => (
              <ServiceCard
                key={index}
                title={service.name}
                image={service.image}
                status={service.isConnected}
                isBlackTheme={isBlackTheme}
                aboutJson={aboutJson}
                serverIp={serverIp}
                setNeedRefresh={setNeedRefresh}
                token={token}
                setModalVisible={setModalVisible}
                setSelectedService={setSelectedService}
                oauth={service.is_oauth}
              />
            ))}
        </View>
        <DeconnectionPopUp
          modalVisible={modalVisible}
          setModalVisible={setModalVisible}
          service={selectedServices}
          token={token}
          serverIp={serverIp}
          setNeedRefresh={setNeedRefresh}
        />
      </View>
    </View>
  );
}
