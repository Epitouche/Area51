import React, { useContext } from 'react';
import {
  View,
  Text,
  StyleSheet,
  ScrollView,
  TouchableOpacity,
} from 'react-native';
import { useNavigation, NavigationProp } from '@react-navigation/native';

import { AppContext } from '../context/AppContext';
import { globalStyles } from '../styles/global_style';
import { IpInput } from '../components';

const workflows = [
  {
    id: 1,
    title: 'GitHub to Email',
    description: 'Send an email when a new github issue is created',
  },
  {
    id: 2,
    title: 'Email to File Saving',
    description: 'Save important files received by email in your drive',
  },
  {
    id: 3,
    title: 'Spotify Notification',
    description:
      'Get a notification from your favorite artist last song directly on the app',
  },
];

const features = [
  {
    icon: '👥',
    title: 'Our Team',
    description:
      'We are a group of 5 passionate students from Epitech Bordeaux united by our ambition to bring innovative ideas to life',
  },
  {
    icon: '💡',
    title: 'The Project',
    description:
      'The goal was to create an AREA website where users like you can create useful workflows',
  },
  {
    icon: '🎯',
    title: 'Our Aim',
    description:
      'Our commitment to excellence and problem-solving drives everything we do. Together, we aim to create something impactful and unforgettable',
  },
];

type RootStackParamList = {
  Register: undefined;
};

const HomeScreen = () => {
  const navigation = useNavigation<NavigationProp<RootStackParamList>>();
  const {
    isBlackTheme,
    aboutJson,
    setAboutJson,
    setServicesConnected,
    serverIp,
    setServerIp,
  } = useContext(AppContext);
  return (
    <View
      style={
        isBlackTheme ? globalStyles.wallpaperBlack : globalStyles.wallpaper
      }>
      <ScrollView
        style={isBlackTheme ? styles.containerBlack : styles.container}>
        <View style={{ width: '90%', alignSelf: 'center', marginBottom: 16 }}>
          <IpInput
            setAboutJson={setAboutJson}
            aboutJson={aboutJson}
            setServicesConnected={setServicesConnected}
            isBlackTheme={isBlackTheme}
            serverIp={serverIp}
            setServerIp={setServerIp}
          />
        </View>
        {/* Hero Section */}
        <View style={styles.hero}>
          <Text
            style={isBlackTheme ? globalStyles.titleBlack : globalStyles.title}>
            Automatize
          </Text>
          <Text style={styles.heroSubtitle}>Everything.</Text>
          <Text style={styles.heroDescription}>
            Make your favorite apps connect with each other and let us
            automatize tasks for you!
          </Text>
          <TouchableOpacity
            style={styles.button}
            onPress={() => navigation.navigate('Register')}>
            <Text style={styles.buttonText}>Start now!</Text>
          </TouchableOpacity>
        </View>

        {/* Workflows */}
        <View style={styles.section}>
          <Text
            style={
              isBlackTheme ? styles.sectionTitleBlack : styles.sectionTitle
            }>
            Make your day easy
          </Text>
          <Text style={styles.sectionSubtitle}>
            Discover new ways to create useful workflows
          </Text>
          {workflows.map(workflow => (
            <View
              key={workflow.id}
              style={isBlackTheme ? styles.cardBlack : styles.card}>
              <Text
                style={isBlackTheme ? styles.cardTitleBlack : styles.cardTitle}>
                {workflow.title}
              </Text>
              <Text style={styles.cardDescription}>{workflow.description}</Text>
            </View>
          ))}
        </View>

        {/* About Us */}
        <View style={styles.section}>
          <Text
            style={
              isBlackTheme ? styles.sectionTitleBlack : styles.sectionTitle
            }>
            About Us
          </Text>
          <Text style={styles.sectionSubtitle}>
            We're on a mission to make workflow automation accessible to
            everyone
          </Text>
          {features.map((feature, index) => (
            <View
              key={index}
              style={isBlackTheme ? styles.cardBlack : styles.card}>
              <Text
                style={isBlackTheme ? styles.cardTitleBlack : styles.cardTitle}>
                {feature.icon} {feature.title}
              </Text>
              <Text style={styles.cardDescription}>{feature.description}</Text>
            </View>
          ))}
          <Text style={styles.link}>
            You can see the documentation on our github repository to learn more
            about us!
          </Text>
        </View>
      </ScrollView>
    </View>
  );
};

const styles = StyleSheet.create({
  container: { flex: 1, backgroundColor: '#F9FAFB', padding: 16 },
  hero: { alignItems: 'center', marginBottom: 32 },
  heroSubtitle: { fontSize: 36, fontWeight: 'bold', color: '#B454FD' },
  heroDescription: { textAlign: 'center', fontSize: 16, color: '#6B7280' },
  button: {
    backgroundColor: '#B454FD',
    paddingVertical: 12,
    paddingHorizontal: 24,
    borderRadius: 8,
    marginTop: 16,
  },
  buttonText: { color: '#FFFFFF', fontWeight: 'bold', fontSize: 16 },
  section: { marginBottom: 32 },
  sectionTitle: { fontSize: 24, fontWeight: 'bold', color: '#111827' },
  sectionSubtitle: { fontSize: 16, color: '#6B7280', marginBottom: 16 },
  card: {
    backgroundColor: '#FFFFFF',
    borderRadius: 8,
    padding: 16,
    marginBottom: 16,
    shadowColor: '#000',
    shadowOpacity: 0.1,
    shadowRadius: 4,
    elevation: 2,
  },
  cardTitle: { fontSize: 18, fontWeight: 'bold', color: '#000000' },
  cardDescription: { fontSize: 14, color: '#6B7280', marginTop: 8 },
  link: { textAlign: 'center', fontSize: 14, color: '#6B7280', marginTop: 16 },

  // Black Theme
  containerBlack: { flex: 1, backgroundColor: '#000000', padding: 16 },
  sectionTitleBlack: { fontSize: 24, fontWeight: 'bold', color: '#F7FAFB' },
  cardBlack: {
    backgroundColor: '#222831',
    borderRadius: 8,
    padding: 16,
    marginBottom: 16,
    shadowColor: '#000',
    shadowOpacity: 0.1,
    shadowRadius: 4,
    elevation: 2,
  },
  cardTitleBlack: { fontSize: 18, fontWeight: 'bold', color: '#F7FAFB' },
});

export default HomeScreen;
