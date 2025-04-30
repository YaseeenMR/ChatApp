import React, { useState, useEffect } from 'react';
import { View, Text, StyleSheet, TextInput, Alert } from 'react-native';
import Button from '../../components/Button'; // adjust path if needed
import { useAuth } from '../../context/AuthContext';
import { updateProfile } from '../../services/auth';

const ProfileScreen = () => {
  const { user, logout } = useAuth();
  const [name, setName] = useState(user?.name || '');
  const [password, setPassword] = useState('');

  const handleUpdate = async () => {
    try {
      const updateData: { name?: string; password?: string } = {};
      if (name !== user?.name) updateData.name = name;
      if (password) updateData.password = password;
      
      if (Object.keys(updateData).length > 0) {
        await updateProfile(updateData);
        Alert.alert('Success', 'Profile updated successfully');
        setPassword('');
      }
    } catch (error: any) {
      Alert.alert('Error', error.message);
    }
  };

  return (
    <View style={styles.container}>
      <Text style={styles.title}>Profile</Text>
      <TextInput
        style={styles.input}
        placeholder="Name"
        value={name}
        onChangeText={setName}
      />
      <TextInput
        style={styles.input}
        placeholder="New Password (leave empty to keep current)"
        value={password}
        onChangeText={setPassword}
        secureTextEntry
      />
      <Button title="Update Profile" onPress={handleUpdate} />
      <Button title="Logout" onPress={logout} />
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    padding: 20,
  },
  title: {
    fontSize: 24,
    marginBottom: 20,
    textAlign: 'center',
  },
  input: {
    height: 40,
    borderColor: 'gray',
    borderWidth: 1,
    marginBottom: 20,
    paddingHorizontal: 10,
  },
});

export default ProfileScreen;