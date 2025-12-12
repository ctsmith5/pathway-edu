import React from 'react';

const variantIcons = {
  info: 'ℹ️',
  tip: '💡',
  warning: '⚠️',
  danger: '🚫',
};

const variantLabels = {
  info: 'Info',
  tip: 'Tip',
  warning: 'Warning',
  danger: 'Important',
};

const CalloutBlock = ({ data }) => {
  if (!data?.text) {
    return null;
  }

  const { variant = 'info', text } = data;
  const icon = variantIcons[variant] || variantIcons.info;
  const label = variantLabels[variant] || variantLabels.info;

  return (
    <div className={`callout-block callout-block--${variant}`}>
      <div className="callout-block__icon">{icon}</div>
      <div className="callout-block__content">
        <span className="callout-block__label">{label}</span>
        <p className="callout-block__text">{text}</p>
      </div>
    </div>
  );
};

export default CalloutBlock;

