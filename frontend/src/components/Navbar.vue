<template>
    <header class="navbar">
        <div class="content-container">
            <router-link
                class="navbar-brand"
                :to="{ name: 'home' }">
                eFreedomWall
            </router-link>
        </div>
    </header>
</template>

<script lang="ts" setup>
import { RouterLink } from 'vue-router';
</script>

<style scoped lang="postcss">
.navbar {
    @apply bg-blue-800 text-white py-4;
}

.navbar > .content-container {
    @apply flex justify-center;
}

.navbar .navbar-brand {
    @apply font-bold text-[2rem];
}
</style>