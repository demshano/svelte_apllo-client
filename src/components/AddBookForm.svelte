<!-- AddBookForm.svelte -->

<script>
    import { mutation } from 'svelte-apollo';
    import { ADD_BOOK } from '../queries';
  
    let title = '';
    let author = '';
  
    const addBookMutation = mutation(ADD_BOOK);
  
    async function handleSubmit() {
      try {
        await addBookMutation({ variables: { title, author } });
      } catch (error) {
        console.error('Error adding book:', error.message);
      }
    }
  </script>
  
  <form on:submit|preventDefault="{handleSubmit}">
    <label for="book-title">Title:</label>
    <input type="text" id="book-title" bind:value="{title}" />
  
    <label for="book-author">Author:</label>
    <input type="text" id="book-author" bind:value="{author}" />
  
    <button type="submit">Add Book</button>
  </form>
  