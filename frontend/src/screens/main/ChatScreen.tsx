import React, { useState, useEffect } from 'react';
import { View, Text, StyleSheet, TextInput, FlatList } from 'react-native';
import Button from '../../components/Button'; // adjust path if needed
import { useAuth } from '../../context/AuthContext';
import MessageBubble from '../../components/MessageBubble';

// Mock data for messages
const mockMessages = [
  { id: '1', text: 'Hello there!', sender: 'other', timestamp: new Date() },
  { id: '2', text: 'Hi! How are you?', sender: 'me', timestamp: new Date() },
];

const ChatScreen = ({ navigation }: any) => {
  const [messages, setMessages] = useState(mockMessages);
  const [newMessage, setNewMessage] = useState('');
  const { user, logout } = useAuth();

  const handleSend = () => {
    if (newMessage.trim()) {
      const message = {
        id: Date.now().toString(),
        text: newMessage,
        sender: 'me',
        timestamp: new Date(),
      };
      setMessages([...messages, message]);
      setNewMessage('');
    }
  };

  return (
    <View style={styles.container}>
      <Text style={styles.header}>Welcome, {user?.name}</Text>
      <FlatList
        data={messages}
        keyExtractor={(item) => item.id}
        renderItem={({ item }) => (
          <MessageBubble
            text={item.text}
            isMe={item.sender === 'me'}
            timestamp={item.timestamp}
          />
        )}
        contentContainerStyle={styles.messagesContainer}
      />
      <View style={styles.inputContainer}>
        <TextInput
          style={styles.input}
          value={newMessage}
          onChangeText={setNewMessage}
          placeholder="Type a message"
        />
        <Button title="Send" onPress={handleSend} />
      </View>
      <Button title="Logout" onPress={logout} />
      <Button title="Update Profile" onPress={() => navigation.navigate('Profile')} />
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    padding: 10,
  },
  header: {
    fontSize: 18,
    fontWeight: 'bold',
    marginBottom: 10,
  },
  messagesContainer: {
    paddingBottom: 10,
  },
  inputContainer: {
    flexDirection: 'row',
    alignItems: 'center',
    marginTop: 10,
  },
  input: {
    flex: 1,
    borderColor: 'gray',
    borderWidth: 1,
    borderRadius: 5,
    padding: 10,
    marginRight: 10,
  },
});

export default ChatScreen;