<!-- App.svelte -->

<script>
	import { ApolloClient, InMemoryCache } from '@apollo/client/core';
	import { setClient } from 'svelte-apollo';
	import Books from './components/Books.svelte';

  
	const client = new ApolloClient({
	  uri: 'http://localhost:8080',
	  cache: new InMemoryCache()
	});
  
	setClient(client);
  </script>
  
  <Books />
  