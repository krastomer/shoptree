import React from "react";
import ItemOrderWait from '../../components/carditems/ItemOrderWait';
import ItemOrderConfirm from '../../components/carditems/ItemOrderConfirm';
import ItemOrderDelivery from '../../components/carditems/ItemOrderDelivery';
import ItemOrderDelivered from '../../components/carditems/ItemOrderDelivered';

// components

export default function HeaderStats() {
  return (
    <> 

    <div className="grid w-full grid-cols-4 gap-4">
        <ItemOrderWait />
        <ItemOrderConfirm />
        <ItemOrderDelivery />
        <ItemOrderDelivered />
    </div>

    </>
  );
}
