package messages

// Subscriptions queues
const SUBSCRIPTIONS_E = "etp_subscriptions_payments"
const SUBSCRIPTIONS_BUY_REQUEST_Q = "qtp_subscriptions_buy_request"
const SUBSCRIPTIONS_BUY_SUCCESS_Q = "qtp_subscriptions_buy_success"
const SUBSCRIPTIONS_BUY_FAIL_Q = "qtp_subscriptions_buy_fail"

// Subscriptions routing keys
const SUBSCRIPTIONS_BUY_REQUEST_K = "subscriptions.buy.request"
const SUBSCRIPTIONS_BUY_SUCCESS_K = "subscriptions.buy.success"
const SUBSCRIPTIONS_BUY_FAIL_K = "subscriptions.buy.fail"

// Payments queues
const PAYMENTS_E = "etp_payment"
const PAYMENTS_APPROVED_Q = "qtp_payment_approved"
const PAYMENTS_REFUSED_Q = "qtp_payment_refused"

// Payment routing keys
const PAYMENTS_APPROVED_K = "payment.approved"
const PAYMENTS_REFUSED_K = "payment.refused"
