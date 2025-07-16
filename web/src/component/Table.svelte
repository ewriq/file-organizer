<script>
    import { onMount } from 'svelte';
    import axios from 'axios';
  
    let files = []; // artık dizi
    let loading = true;
    let error = '';
  
    onMount(async () => {
      try {
        const res = await axios.get('http://localhost:3000/api/v1/files/'); // örnek uzantı/kategori
        files = res.data;
      } catch (err) {
        error = err.message || 'Veri alınırken hata oluştu';
      } finally {
        loading = false;
      }
    });
  </script>
  
  <div class="flex items-center justify-center min-h-screen bg-gray-400">
    <div class="w-4/5 border border-indigo-400 shadow-[0_10px_100px_2px_rgba(59,100,246,0.6)] text-white p-8 rounded-lg bg-gray-700">
      
      {#if loading}
        <p>Yükleniyor...</p>
      {:else if error}
        <p class="text-red-500">Hata: {error}</p>
      {:else if files.length > 0}
        <h2 class="text-2xl font-bold mb-6 text-center">Dosyalar ({files.length})</h2>
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
          {#each files as file}
            <div class="bg-gray-100 text-black p-4 rounded shadow hover:shadow-xl transition">
              <h3 class="font-semibold text-lg mb-2">{file.Name}</h3>
              <ul class="text-sm">
                <li><strong>ID:</strong> {file.ID}</li>
                <li><strong>Tür:</strong> {file.Type}</li>
                <li><strong>Boyut:</strong> {file.Size} byte</li>
                <li><strong>Zaman:</strong> {file.Time}</li>
                <li><strong>Yol:</strong> <span class="break-words">{file.Path}</span></li>
              </ul>
            </div>
          {/each}
        </div>
      {:else}
        <p>Dosya bulunamadı.</p>
      {/if}
  
    </div>
  </div>
  