<template>
  <div class="bumble-clone">
    <div class="sidebar">
      <div class="user-info">
        <img :src="currentUser.avatar" :alt="currentUser.name" class="avatar">
        <span>{{ currentUser.name }}</span>
      </div>
      <div class="match-queue">
        <h3>Match Queue</h3>
        <div class="spotlight">
          <span class="icon">🌟</span>
          <p>Spotlight is the easiest way to up your odds of a match!</p>
          <a href="#">Learn More</a>
        </div>
      </div>
      <div class="conversations">
        <h3>Conversations (Recent)</h3>
        <ul>
          <li v-for="conversation in conversations" :key="conversation.id">
            <img :src="conversation.avatar" :alt="conversation.name" class="avatar">
            <div>
              <span class="name">{{ conversation.name }}</span>
              <span class="message">{{ conversation.lastMessage }}</span>
            </div>
          </li>
        </ul>
      </div>
    </div>
    <div class="main-content">
      <div class="filters">
        <button>Filters</button>
      </div>
      <div class="profile-card">
        <img :src="currentProfile.image" :alt="currentProfile.name" class="profile-image">
        <div class="profile-info">
          <h2>{{ currentProfile.name }}, {{ currentProfile.age }}</h2>
          <p>{{ currentProfile.occupation }}</p>
          <div class="actions">
            <button class="action-button block">✕</button>
            <button class="action-button favorite">⭐</button>
            <button class="action-button like">✓</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue'

export default {
  name: 'BumbleClone',
  setup() {
    const currentUser = ref({
      name: 'James',
      avatar: '/path/to/james-avatar.jpg'
    })

    const conversations = ref([])
    const currentProfile = ref({})

    const fetchConversations = async () => {
      // Fetch conversations from API
      // This is a placeholder. You'll need to implement the actual API call.
      conversations.value = [
        { id: 1, name: 'Jhe', avatar: '/path/to/jhe-avatar.jpg', lastMessage: 'Conversation expired on 06.14.24' },
        { id: 2, name: 'Ping', avatar: '/path/to/ping-avatar.jpg', lastMessage: 'If I can help you practice your english, let ...' },
        // ... more conversations
      ]
    }

    const fetchCurrentProfile = async () => {
      // Fetch current profile from API
      // This is a placeholder. You'll need to implement the actual API call.
      currentProfile.value = {
        name: 'Sofia',
        age: 52,
        image: '/path/to/sofia-image.jpg',
        occupation: 'Healthcare at Healthcare'
      }
    }

    onMounted(() => {
      fetchConversations()
      fetchCurrentProfile()
    })

    return {
      currentUser,
      conversations,
      currentProfile
    }
  }
}
</script>

<style scoped>
.bumble-clone {
  display: flex;
  font-family: Arial, sans-serif;
}

.sidebar {
  width: 300px;
  padding: 20px;
  border-right: 1px solid #e0e0e0;
}

.main-content {
  flex-grow: 1;
  padding: 20px;
}

.avatar {
  width: 50px;
  height: 50px;
  border-radius: 50%;
}

.profile-image {
  width: 100%;
  max-height: 500px;
  object-fit: cover;
}

.actions {
  display: flex;
  justify-content: space-around;
  margin-top: 20px;
}

.action-button {
  width: 60px;
  height: 60px;
  border-radius: 50%;
  border: none;
  font-size: 24px;
  cursor: pointer;
}

.block { background-color: #f0f0f0; }
.favorite { background-color: #ffd700; }
.like { background-color: #00ff00; }
</style>
