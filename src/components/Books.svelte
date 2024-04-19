<!-- Books.svelte -->

<script>
    import { query } from 'svelte-apollo';
    import { GET_BOOKS } from '../queries';
  
    const { data, loading, error } = query(GET_BOOKS);
  </script>
  
  {#if loading}
    <p>Loading...</p>
  {:else if error}
    <p>Error: {error.message}</p>
  {:else}
    <ul>
      {#each data.books as book (book.id)}
        <li>{book.title} by {book.author.name}</li>
      {/each}
    </ul>
  {/if}
  