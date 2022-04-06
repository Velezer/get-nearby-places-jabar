window.onload = function() {
  //<editor-fold desc="Changeable Configuration Block">
  let host = window.location.origin
  // the following lines will be replaced by docker/configurator, when it runs in a docker-container
  window.ui = SwaggerUIBundle({
    url: `${host}/docs/spec/jcc-openapi-spec.yaml`,
    dom_id: '#swagger-ui',
    deepLinking: true,
    presets: [
      SwaggerUIBundle.presets.apis,
      SwaggerUIStandalonePreset
    ],
    plugins: [
      SwaggerUIBundle.plugins.DownloadUrl
    ],
    layout: "StandaloneLayout"
  });

  window.ui.api.setHost(host)

  //</editor-fold>
};
