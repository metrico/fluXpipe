window.$docsify = {
  name: 'Flux Community',
  repo: 'https://github.com/InfluxCommunity/flux',
  loadSidebar: true,
  auto2top: true,
  coverpage: true,
  basePath: '.', // Adjust this base path as needed
  alias: {
    '/_sidebar.md': '_sidebar.md',
  },
};

// Initialize Docsify
document.addEventListener('DOMContentLoaded', function () {
  docsify.init();
});

