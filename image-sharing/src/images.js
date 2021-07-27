const context = require.context("@/assets/images", false, /\.jpg$/);
const images = context.keys().map(context);

export default images